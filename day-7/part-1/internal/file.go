package internal

type File struct {
	Name string
	Size int
}

type Folder struct {
	Name      string
	TotalSize int
	Files     []*File
	Folders   map[string]*Folder
}

type FileSystem struct {
	RootFolder *Folder
}

func (f *FileSystem) AddFolder(cwd []string, folderNameToAdd string) {
	currentFolder := f.RootFolder
	for _, s := range cwd {
		currentFolder = currentFolder.Folders[s]
	}
	currentFolder.AddFolder(&Folder{
		Name:      folderNameToAdd,
		TotalSize: 0,
		Files:     []*File{},
		Folders:   make(map[string]*Folder),
	})
}

func (f *FileSystem) AddFile(cwd []string, fileToAdd *File) {
	currentFolder := f.RootFolder
	currentFolder.TotalSize += fileToAdd.Size

	for _, s := range cwd {
		currentFolder = currentFolder.Folders[s]
		currentFolder.TotalSize += fileToAdd.Size
	}
	currentFolder.AddFile(fileToAdd)
}

func (f *Folder) AddFile(theFile *File) {
	f.Files = append(f.Files, theFile)
}

func (f *Folder) AddFolder(theFolder *Folder) {
	f.Folders[theFolder.Name] = theFolder
}
