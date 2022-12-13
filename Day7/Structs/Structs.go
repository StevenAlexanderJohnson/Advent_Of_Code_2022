package structs

type Directory struct {
	Name             string
	Parent_Directory *Directory
	Files            map[string]File
	Directories      map[string]Directory
}

type DirectoryNotFound struct{}

func (d *DirectoryNotFound) Error() string {
	return "Directory not found"
}

type File struct {
	Name string
	Size uint64
}

func (d *Directory) Add_File(name string, size uint64) {
	d.Files[name] = File{Name: name, Size: size}
}

func (d *Directory) Add_Directory(name string) {
	d.Directories[name] = Directory{Name: name, Parent_Directory: d, Files: make(map[string]File), Directories: make(map[string]Directory)}
}

// return the directory if it exists, otherwise return an error
func (d *Directory) CD(name string) (*Directory, error) {
	if name == ".." && d.Parent_Directory != nil {
		return d.Parent_Directory, nil
	} else if directory, ok := d.Directories[name]; ok {
		return &directory, nil
	} else {
		return nil, &DirectoryNotFound{}
	}
}

func (d *Directory) Get_Size() uint64 {
	var output uint64 = 0
	var current_directory_size uint64 = 0
	for _, file := range d.Files {
		current_directory_size += file.Size
	}

	for _, directory := range d.Directories {
		temp_size := directory.Get_Size()
		if temp_size <= 100000 {
			output += temp_size
		}
		current_directory_size += temp_size
	}
	if current_directory_size <= 100000 {
		output += current_directory_size
	}
	return output
}
