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
rotate(angle int)
flipX()
flipY()
crop(x0, y0, x1, y1 int)
resize(width, height int)
brightness(percentage int)
contrast(percentage int)
saturation(percentage int)
pixelate(size int)
```

### Errors:
1. Misspelled command (ex: opn, instead of open)
2. Missing token
3. Wrong file name
4. Using a non delcared variable
5. Invalid arguments in actions (ex: string instead of float or int)
6. Invalid image type or non image type.