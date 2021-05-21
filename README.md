### Build and Run
``` bash
go mod vendor
go build
./img.exe input.txt
``` 

### Program Example

``` bash
-- Declarations --
-- Image
img = open("capture.png");
-- Folder
folder[] = open("images"); 

-- Actions --
img.rotate(180);
folder.pixelate(10);

-- Exports --
img.save("new_image_1.png");
folder.save("export");
```

### Actions available:
``` bash
rotate(angle float)
flipX()
flipY()
crop(x0, y0, x1, y1 float)
resize(width, height float)
brightness(percentage float)
contrast(percentage float)
saturation(percentage float)
pixelate(size int)
```

### Errors:
1. Misspelled command (ex: opn, instead of open)
2. Missing token
2. Wrong file name
3. Using a non delcared variable
4. Invalid arguments in actions (ex: string instead of float or int)
