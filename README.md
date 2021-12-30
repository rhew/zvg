# Test communication with ZVG

```
sudo su -                                               # access parallel port address
export ZVGPORT="P3BC M4"                                # update to match your address and monitor type
export LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH  # update to include the zvg linux library path
go run ./cmd/test.go                                    # check for non-zero version number
```
