### Build and Run
``` bash
go build
./img.exe input.img
``` 

### Program Structure

``` bash
--- Declarations ---
img1 = open("image1.png");
img2 = open("image2.png");

--- Actions ---
img1.rotate(180);
img2.pixelate(15);

--- Exports ---
img1.save("new_image_1.png");
img2.save("new_image_2.png");
```

**Note:** Comments are not defined by grammar, yet.

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