package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

type Contact struct {
	ID        bson.ObjectId `bson:"_id"` // Saved to mongo as "_id".
	Email     string
	FirstName string
	LastName  string
}

func main() {
	// sess is a reference to an mgo.Session object,
	// which has several methods for manipulating the database.
	// I am using Windows Home which doesn't have a native hypervisor.
	// So my IP address for Dial is the IP address of my Linux virtual machine/
	sess, err := mgo.Dial("192.168.99.100")
	if err != nil {
		fmt.Printf("error dialing mongo: %v\n", err)
	} else {
		fmt.Printf("connected successfully!\n")
	}

	c := &Contact{
		ID:        bson.NewObjectId(),
		Email:     "test@test.com",
		FirstName: "Test",
		LastName:  "Tester",
	}

	// Get a reference to the "contacts" collection
	// in the "demo" database.
	coll := sess.DB("demo").C("contacts")

	// Insert struct into that collection.
	if err := coll.Insert(c); err != nil {
		fmt.Printf("error inserting document: %v\n", err)
	} else {
		// An objectId is actually a 12-byte binary value,
		// so we need to print it in a human-readable form.
		fmt.Printf("inserted document with ID %s\n", c.ID.Hex())
	}

	// Create an empty slice of
	// pointers to Contact structs.
	contacts := []*Contact{}

	// Find all documents where the lastname property contains "Tester"
	// and decode them into the slice
	// (using same `coll` variable from earlier code snippet)
	coll.Find(bson.M{"lastname": "Tester"}).All(&contacts)

	// Iterate over the slice, printing the data to std out.
	for _, c := range contacts {
		fmt.Printf("%s: %s, %s, %s\n", c.ID.Hex(), c.Email, c.FirstName, c.LastName)
	}

	/*
		If we don't know how much data we will be getting
		and not sure if the results will fit into memory at once,
		we can use .Iter() to process data one at a time.

		// Declare just one empty contact struct
		// to receive one document at a time.
		c = &Contact{}

		// Create an iterator for the found documents.
		iter := coll.Find(bson.M{"lastname": "Tester"}).Iter()

		// For as long as we get another document from the iterator...
		for iter.Next(c) {
			// Print the data to std out.
			fmt.Printf("%s: %s, %s, %s\n", c.ID.Hex(), c.Email, c.FirstName, c.LastName)
		}

		// Report any errors that occurred.
		if err := iter.Err(); err != nil {
			fmt.Printf("error iterating found documents: %v\n", err)
		}
	*/

	/*
		Finding document by its unique ID.

		c = &Contact{}

		// Assuming a variable `id`
		// set to the ObjectId you want to find.
		coll.FindId(id).One(c)
	*/

	id := c.ID //...bson ObjectId of the document you want to update.

	// Create a map with the properties you want to update.
	// You can alternatively use a struct here instead.
	updates := bson.M{"firstname": "UPDATED Test"}

	// Send an update document with a `$set` property set to the updates map.
	// The value of $set is another map.
	if err := coll.UpdateId(id, bson.M{"$set": updates}); err != nil {
		fmt.Printf("error updating document: %v\n", err)
	}

	fmt.Println("document updated")

	// Assuming `id` has been set to a bson.ObjectId...
	if err := coll.RemoveId(id); err != nil {
		fmt.Printf("error deleting document: %v\n", err)
	}

	fmt.Println("document deleted")
}
