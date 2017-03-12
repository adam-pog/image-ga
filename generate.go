package main

import (
    "github.com/llgcode/draw2d/draw2dimg"
    "image"
    "image/color"
    "math/rand"
    "time"
    "fmt"
    "strconv"
    "os"
)

const Xmax = 1920
const Ymax = 1080
const PopSize = 10
const GenePoolSize = 10

type Triangle struct {
    x1, y1, x2, y2, x3, y3 float64
    color color.RGBA
}

type Chromosome struct {
    triangles [GenePoolSize]Triangle
    fitness int
    gc *draw2dimg.GraphicContext
    genImage *image.RGBA
}

func main() {
    // Initialize the graphic context on an RGBA image
    rand.Seed(time.Now().Unix())
    //filename := "/home/npzd/Pictures/generated/image" + strconv.FormatInt(time.Now().Unix(), 10) + ".png"
    //draw2dimg.SaveToPngFile(filename, gen1[0].genImage)

    fmt.Println("--- Initializing Population ---")
    gen1 := initializePopulation()

    fmt.Println("Drawing Chromosomes")
    for i := 0; i < PopSize; i++ {
        fmt.Println("Drawing Chromosome " + strconv.FormatInt(int64(i), 10))
        gc := gen1[i].gc
        drawTriangles(gc, gen1[i].triangles)
    }
    infile, _ := os.Open("/home/npzd/Pictures/owsombra.png")
    sourceImg, _, _ := image.Decode(infile)

    calculateFitness(&gen1, sourceImg)

    fmt.Println("done")
}

func calculateFitness(chromosomes *[PopSize]Chromosome, sourceImage image.Image) {
    for i := 0; i < PopSize; i++ {
        fmt.Println("Comparing Chromosome " + strconv.FormatInt(int64(i), 10))
        diff := 0
        for x := 0; x < Xmax; x++ {
            for y := 0; y < Ymax; y++ {
                r, g, b, a := chromosomes[i].genImage.At(x, y).RGBA()
                r2, g2, b2, a2 := sourceImage.At(x, y).RGBA()
                diff += int(r) + int(r2)
                diff += int(g) + int(g2)
                diff += int(b) + int(b2)
                diff += int(a) + int(a2)
            }
        }
        chromosomes[i].fitness = diff
    }

}

func drawTriangles(gc *draw2dimg.GraphicContext, triangles [GenePoolSize]Triangle) {
    for i := 0; i < GenePoolSize; i++ {
        gc.SetFillColor(triangles[i].color)
        gc.MoveTo(triangles[i].x1, triangles[i].y1)
        gc.LineTo(triangles[i].x2, triangles[i].y2)
        gc.LineTo(triangles[i].x3, triangles[i].y3)
        gc.Close()
        gc.Fill()
    }
}

func initializePopulation() (population [PopSize]Chromosome) {
    for i := 0; i < PopSize; i++ {
        population[i] = generateChromosome()
    }
    return
}

func generateChromosome() (chromosome Chromosome) {
    chromosome.genImage = image.NewRGBA(image.Rect(0, 0, Xmax, Ymax))
    chromosome.gc = draw2dimg.NewGraphicContext(chromosome.genImage)

    for i := 0; i < GenePoolSize; i++ {
        chromosome.triangles[i] = generateTriangle()
    }
    return
}

func generateTriangle() Triangle {
    return Triangle{x1: random(0, Xmax),
                    y1: random(0, Ymax),
                    x2: random(0, Xmax),
                    y2: random(0, Ymax),
                    x3: random(0, Xmax),
                    y3: random(0, Ymax),
                    color: generateColor()}
}

func generateColor() color.RGBA {
    return color.RGBA{uint8(random(0, 255)), uint8(random(0, 255)), uint8(random(0, 255)), uint8(random(0, 255))}
}

func random(min, max int) float64 {
    return float64(rand.Intn(max - min) + min)
}
