package generator

import "os"

var PATH, _ = os.Getwd()

/*
	BIN path with compiled html pubs
*/
var PATH_BIN = PATH + "/bin"

/*
	Folder with all DRAFTS 
*/
var PATH_DRAFTS = PATH + "/drafts"

/*
	Folder with TEMPLATES (Pub & Pub List)
*/
var PATH_TEMPLATES = PATH + "/templates"