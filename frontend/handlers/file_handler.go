package handlers

import (
	"context"
	"fmt"
	"html/template"
	"main/helpers"
	pb "main/proto"
	"net/http"
)

type PageData struct {
	Objects []*pb.Object
}

func FileHandler(w http.ResponseWriter, r *http.Request, directory string) {
	tmpl := template.Must(template.ParseFiles("templates/files.html"))

	clientRetriever := helpers.GetRetrieverClient()

	//TODO: Make the bucket name dynamic!
	structureRequest := &pb.StructureRequest{Bucket: "veit", Directory: "src/test/java/backend/"}
	structure, err := clientRetriever.GetStructure(context.Background(), structureRequest)
	if err != nil {
		fmt.Println("Failed to get structure: ", err)
		return
	}

	//TODO: Add handling for empty directory

	for i := 0; i < len(structure.Object); i++ {
		fmt.Println("Name: ", structure.Object[i].Name, "Type: ", structure.Object[i].Type)
	}

	data := PageData{Objects: structure.Object}
	tmpl.Execute(w, data)
}