package controllers
//
//import (
//	"time"
//
//	"github.com/GeertJohan/go.rice/embedded"
//)
//
//func init() {
//
//	// define files
//	file2 := &embedded.EmbeddedFile{
//		Filename:    "generate.html",
//		FileModTime: time.Unix(1551322309, 0),
//
//		Content: string("<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"UTF-8\">\n  <title>Title</title>\n</head>\n<body>\n{{if .Success}}\n<h1>Thanks for your message!</h1>\n{{else}}\n<h1>Contact</h1>\n<form method=\"POST\">\n  <label>Email:</label><br />\n  <input type=\"text\" name=\"email\"><br />\n  <label>Subject:</label><br />\n  <input type=\"text\" name=\"subject\"><br />\n  <label>Message:</label><br />\n  <textarea name=\"message\"></textarea><br />\n  <input type=\"submit\">\n</form>\n{{end}}\n</body>\n</html>"),
//	}
//
//	file3 := &embedded.EmbeddedFile{
//		Filename:    "home.html",
//		FileModTime: time.Unix(1551410754, 0),
//
//		Content: string("<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"UTF-8\">\n  <title>\n   </title>\n</head>\n<body>\n<h1> This is Generate Page ! </h1>\n  <ul>\n  {{range .}}\n    <li class=\"done\">{{.Name}}</li>\n  {{end}}\n  </ul>\n</body>\n</html>"),
//	}
//
//	// define dirs
//	dir1 := &embedded.EmbeddedDir{
//		Filename:   "",
//		DirModTime: time.Unix(1551410754, 0),
//		ChildFiles: []*embedded.EmbeddedFile{
//			file2, // "generate.html"
//			file3, // "home.html"
//		},
//	}
//
//	// link ChildDirs
//	dir1.ChildDirs = []*embedded.EmbeddedDir{}
//
//	// register embeddedBox
//	embedded.RegisterEmbeddedBox(`../html`, &embedded.EmbeddedBox{
//		Name: `../html`,
//		Time: time.Unix(1551410754, 0),
//		Dirs: map[string]*embedded.EmbeddedDir{
//			"": dir1,
//		},
//		Files: map[string]*embedded.EmbeddedFile{
//			"generate.html": file2,
//			"home.html":     file3,
//		},
//	})
//
//}
