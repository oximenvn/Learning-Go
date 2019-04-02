package main
 
import (
    "fmt"
    "os"
    "image"
    "image/jpeg"
    "image/color"
    "image/png"
)
 

 
func main() {
    //Open file
    file, err := os.Open("image.jpg")
    defer file.Close()

    if err != nil {
        fmt.Println(err)
        return
    }

    //Decode image
    img, err := jpeg.Decode(file)
    if err != nil {
        fmt.Println(err)
        return
    }

    //Calculate
    red, green, blue, gray :=histogram_cal(img)

    //Save image
    out, _ := os.Create("histogram-red.png")
    out1, _ := os.Create("histogram-green.png")
    out2, _ := os.Create("histogram-blue.png")
    out3, _ := os.Create("histogram-gray.png")

    defer out.Close()
    defer out1.Close()
    defer out2.Close()
    defer out3.Close()
    png.Encode(out,red)
    png.Encode(out1,green)
    png.Encode(out2,blue)
    png.Encode(out3,gray)
    //

}

/*
* Calculate histogram of picture
* @arg:
* @return: image histogram of red, green, blue and gray
*/
func histogram_cal(img image.Image) (*image.RGBA, *image.RGBA, *image.RGBA, *image.RGBA) {
        //Create histogram image
        width := 256
        height := 256
    
        upLeft := image.Point{0, 0}
        lowRight := image.Point{width, height}
    
        red := image.NewRGBA(image.Rectangle{upLeft, lowRight})
        green := image.NewRGBA(image.Rectangle{upLeft, lowRight})
        blue := image.NewRGBA(image.Rectangle{upLeft, lowRight})
        gray := image.NewRGBA(image.Rectangle{upLeft, lowRight})

        //Create histogram array
        var histogram  [256][4]int

        //Calculate
        bounds := img.Bounds()
        for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
            for x := bounds.Min.X; x < bounds.Max.X; x++ {
                r, g, b, _ := img.At(x, y).RGBA()
                // rgb to gray
                gray := int(0.299 * float64(r) + 0.587 * float64(g) + 0.114 * float64(b))
                // A color's RGBA method returns values in the range [0, 65535].
                // Shifting by 8 reduces this to the range [0, 255].
                histogram[r>>8][0]++
                histogram[g>>8][1]++
                histogram[b>>8][2]++
                histogram[gray>>8][3]++
            }
        }

        //Normally
        max := get_max(histogram)
        fmt.Println(max)
        for i := range histogram{
            histogram[i][0]= histogram[i][0]*255/max
            histogram[i][1]= histogram[i][1]*255/max
            histogram[i][2]= histogram[i][2]*255/max
            histogram[i][3]= histogram[i][3]*255/max
        }

        //Draw 
        for x:=0; x < 255; x++{
            draw_line(red, x, histogram[x][0], color.RGBA{200, 0, 0, 255})
            draw_line(green, x, histogram[x][1], color.RGBA{0, 200, 0, 255})
            draw_line(blue, x, histogram[x][2], color.RGBA{0, 0, 200, 255})
            draw_line(gray, x, histogram[x][3], color.RGBA{0, 0, 0, 255})
        }
        return red, green, blue, gray
}

func draw_line(img *image.RGBA, position int, length int, color color.RGBA) {
    bounds := (*img).Bounds()
    if bounds.Min.Y <= position && position <= bounds.Max.Y {
        for i := 255; i > 255-length; i-- {
            (*img).Set(position, i, color)
        }
    }
}

func get_max(array [256][4]int) int {
    max:=0
    for i := range array {
        item := array[i]
        local_max := item[0] 

        if local_max < item[1] {
            local_max = item[1]
        }

        if local_max < item[2] {
            local_max = item[2]
        }

        if local_max < item[3] {
            local_max = item[3]
        }

        if max < local_max {
            max = local_max
        }
        /*if max < item{
            max = item
        }*/
    }

    return max
}


