package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/jcelliott/lumber"
)

const Version = "1.0.0"

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}

	Driver struct {
		mutex   sync.Mutex
		mutexes map[string]*sync.Mutex
		dir     string
		log     Logger
	}
)

type Options struct {
	Logger
}

func New(dir string, options *Options) (*Driver, error) {
	dir = filepath.Clean(dir)

	opts := Options{}

	if options != nil {
		opts = *options
	}

	if opts.Logger == nil {
		opts.Logger = lumber.NewConsoleLogger((lumber.INFO))
	}

	driver := Driver{
		dir:     dir,
		mutexes: make(map[string]*sync.Mutex),
		log:     opts.Logger,
	}

	if _, err := os.Stat(dir); err == nil {
		opts.Logger.Debug("Using '%s' (database already exists)\n", dir)
		return &driver, nil
	}

	opts.Logger.Debug("Creating the database at '%s'...\n", dir)
	return &driver, os.MkdirAll(dir, 0755)
}

func (d *Driver) Write(collection, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("Missing collection - no place to save record!")
	}

	if resource == "" {
		return fmt.Errorf("Missing resource - unable to save record (no name)!")
	}

	mutex := d.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, collection)
	fnlPath := filepath.Join(dir, resource+".json")
	tmpPath := fnlPath + ".tmp"

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}

	b = append(b, byte('\n'))

	if err := os.WriteFile(tmpPath, b, 0644); err != nil {
		return err
	}

	return os.Rename(tmpPath, fnlPath)
}

func (d *Driver) Read(collection, resource string, v interface{}) error {

	if collection == "" {
		return fmt.Errorf("Missing collection - unable to read!")
	}

	if resource == "" {
		return fmt.Errorf("Missing resource - unable to read record (no name)!")
	}

	record := filepath.Join(d.dir, collection, resource)

	if _, err := stat(record); err != nil {
		return err
	}

	b, err := os.ReadFile(record + ".json")
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &v)
}

func (d *Driver) ReadAll(collection string) ([]string, error) {

	if collection == "" {
		return nil, fmt.Errorf("Missing collection - unable to read")
	}
	dir := filepath.Join(d.dir, collection)

	if _, err := stat(dir); err != nil {
		return nil, err
	}

	files, _ := os.ReadDir(dir)

	var records []string

	for _, file := range files {
		b, err := os.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}

		records = append(records, string(b))
	}
	return records, nil
}

func (d *Driver) Delete(collection, resource string) error {

	path := filepath.Join(collection, resource)
	mutex := d.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, path)

	switch fi, err := stat(dir); {
	case fi == nil, err != nil:
		return fmt.Errorf("unable to find file or directory named %v\n", path)

	case fi.Mode().IsDir():
		return os.RemoveAll(dir)

	case fi.Mode().IsRegular():
		return os.RemoveAll(dir + ".json")
	}
	return nil
}

func (d *Driver) getOrCreateMutex(collection string) *sync.Mutex {

	d.mutex.Lock()
	defer d.mutex.Unlock()
	m, ok := d.mutexes[collection]

	if !ok {
		m = &sync.Mutex{}
		d.mutexes[collection] = m
	}

	return m
}

func stat(path string) (fi os.FileInfo, err error) {
	if fi, err = os.Stat(path); os.IsNotExist(err) {
		fi, err = os.Stat(path + ".json")
	}
	return
}

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {
	dir := "./"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- KnoxDB Operations Menu ---")
		fmt.Println("1. Add initial dummy employees")
		fmt.Println("2. Write a new user")
		fmt.Println("3. Read a user")
		fmt.Println("4. Read all users")
		fmt.Println("5. Delete a user")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)

		switch choiceStr {
		case "1":
			employees := []User{
				{"Knox", "11", "1234567890", "Knoxverse", Address{"Jalgaon", "Maharashtra", "India", "425001"}},
				{"Luffy", "21", "246801379", "StrawHats", Address{"East-Blue", "GrandLine", "One Piece", "200001"}},
				{"Tonystark", "3000", "0070070070", "Stark Industries", Address{"New York", "unknown", "USA", "123456"}},
			}
			for _, value := range employees {
				db.Write("users", value.Name, value)
			}
			fmt.Println("Dummy employees added successfully.")

		case "2":
			fmt.Print("Enter Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			
			fmt.Print("Enter Age: ")
			ageStr, _ := reader.ReadString('\n')
			ageStr = strings.TrimSpace(ageStr)
			
			newUser := User{
				Name: name, 
				Age: json.Number(ageStr),
			}
			
			err := db.Write("users", newUser.Name, newUser)
			if err != nil {
				fmt.Println("Error writing user:", err)
			} else {
				fmt.Println("User", newUser.Name, "added successfully.")
			}

		case "3":
			fmt.Print("Enter name of the user to read: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			var userFound User
			err := db.Read("users", name, &userFound)
			if err != nil {
				fmt.Println("Error reading user:", err)
			} else {
				fmt.Printf("User details: %+v\n", userFound)
			}

		case "4":
			records, err := db.ReadAll("users")
			if err != nil {
				fmt.Println("Error reading all users:", err)
			} else {
				fmt.Println("All Users in database:")
				allusers := []User{}
				for _, f := range records {
					var employeeFound User
					if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
						fmt.Println("Error unmarshaling:", err)
						continue
					}
					allusers = append(allusers, employeeFound)
				}
				for _, u := range allusers {
					fmt.Printf("- %+v\n", u)
				}
			}

		case "5":
			fmt.Print("Enter name of the user to delete: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			
			err := db.Delete("users", name)
			if err != nil {
				fmt.Println("Error deleting user:", err)
			} else {
				fmt.Println("User", name, "deleted successfully.")
			}

		case "6":
			fmt.Println("Exiting process.")
			return

		default:
			fmt.Println("Invalid choice, please select an option from 1 to 6.")
		}
	}
}
