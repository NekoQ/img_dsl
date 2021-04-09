### Build and Run
``` bash
go build
./img.exe input.txt
``` 

### Program Example

``` bash
img = open("image.png");   
img.rotate(180);
img.save("new_image.png");
```

**Open** and **save** actions are mandatory

### Actions available:
``` bash
rotate(angle float)
flipX()
flipY()
crop(x0, y0, x1, y1 float)
```