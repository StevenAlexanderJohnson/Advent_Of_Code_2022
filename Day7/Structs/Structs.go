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

type FileAlreadyExists struct{}

func (f *FileAlreadyExists) Error() string {
	return "File already exists"
}

type DirectoryAlreadyExists struct{}

func (d *DirectoryAlreadyExists) Error() string {
	return "Directory already exists"
}

type File struct {
	Name string
	Size uint64
}

func (d *Directory) Add_File(name string, size uint64) error {
	_, ok := d.Files[name]
	if ok {
		return &FileAlreadyExists{}
	}
	d.Files[name] = File{Name: name, Size: size}
	return nil
}

func (d *Directory) Add_Directory(name string) error {
	_, ok := d.Directories[name]
	if ok {
		return &DirectoryAlreadyExists{}
	}
	d.Directories[name] = Directory{Name: name, Parent_Directory: d, Files: make(map[string]File), Directories: make(map[string]Directory)}
	return nil
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

func (d *Directory) Get_Size() (uint64, uint64) {
	var output uint64 = 0
	var current_directory_size uint64 = 0
	for _, file := range d.Files {
		current_directory_size += file.Size
	}
	for _, directory := range d.Directories {
		temp_size, previous_size := directory.Get_Size()
		// if temp_size <= 100000 {
		output += temp_size
		// }
		current_directory_size += previous_size
	}
	if current_directory_size <= 100000 {
		output += current_directory_size
	}
	return output, current_directory_size
}
