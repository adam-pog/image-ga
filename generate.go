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
const GenePoolSize = 50

//configuration for producing next gen
const TournamentK = 4
const CrossoverRate = 0.6
const ColorMutationRate = 0.05
const ShapeMutationRate = 0.05

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

    //fmt.Println("Drawing Chromosomes")
    drawChromosomes(gen1)

    infile, _ := os.Open("/home/npzd/Pictures/owsombra.png")
    sourceImg, _, _ := image.Decode(infile)

    calculateFitness(&gen1, sourceImg)

    newestGen := gen1
    for p := 0; p < 1000; p++ {
        newest := createNextGeneration(newestGen)
        fmt.Println("Running generation " + strconv.FormatInt(int64(p), 10))
        drawChromosomes(newest)
        calculateFitness(&newest, sourceImg)
        newestGen = newest
    }
    fmt.Println("done")
}

func calculateFitness(chromosomes *[PopSize]Chromosome, sourceImage image.Image) {
    for i := 0; i < PopSize; i++ {
        //fmt.Println("Comparing Chromosome " + strconv.FormatInt(int64(i), 10))
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

func drawChromosomes(currentGen [PopSize] Chromosome) {
    for i := 0; i < PopSize; i++ {
        //fmt.Println("Drawing Chromosome " + strconv.FormatInt(int64(i), 10))
        gc := currentGen[i].gc
        drawTriangles(gc, currentGen[i].triangles)
    }
}

func createNextGeneration(currentGen [PopSize]Chromosome) (nextGen [PopSize]Chromosome){
    //fmt.Println("Creating generation " + strconv.FormatInt(genNum, 10))
    exportFittestChromosome(currentGen)
    for i := 0; i < PopSize; i++ {
        //parent1 := selectParent(currentGen)
        //parent2 := selectParent(currentGen)


        // if CrossoverRate > rand.Float64() {
        //     nextGen[i] = crossover(parent1, parent2)
        // } else{
        //     if parent1.fitness > parent2.fitness {
        //         nextGen[i] = parent1
        //     }else {
        //         nextGen[i] = parent2
        //     }
        // }
        fittest := currentGen[0]
        for i := 1; i < PopSize; i++ {
            if currentGen[i].fitness < fittest.fitness {
                fittest = currentGen[i]
            }
        }

        nextGen[i] = makeCH(fittest)

        // if parent1.fitness < parent2.fitness {
        //     nextGen[i] = makeCH(parent1)
        // }else {
        //     nextGen[i] = makeCH(parent2)
        // }


        mutateColor(&nextGen[i].triangles)
        mutateShape(&nextGen[i].triangles)
    }
    return nextGen
}

func makeCH(chrom Chromosome) (chromosome Chromosome) {
    chromosome.genImage = image.NewRGBA(image.Rect(0, 0, Xmax, Ymax))
    chromosome.gc = draw2dimg.NewGraphicContext(chromosome.genImage)
    chromosome.triangles = chrom.triangles
    return
}

func mutateColor(triangles *[GenePoolSize]Triangle) {
    for i := 0; i < GenePoolSize; i++ {
        if ColorMutationRate > rand.Float64() {
            triangles[i].color = generateColor()
        }
    }
}

func mutateShape(triangles *[GenePoolSize]Triangle) {
    for i := 0; i < GenePoolSize; i++ {
        if ShapeMutationRate > rand.Float64() {
            triangles[i].x1 = random(0, Xmax)
            triangles[i].y1 = random(0, Ymax)
            triangles[i].x2 = random(0, Xmax)
            triangles[i].y2 = random(0, Ymax)
            triangles[i].x3 = random(0, Xmax)
            triangles[i].y3 = random(0, Ymax)
        }
    }
}

func crossover(chrom1, chrom2 Chromosome) (finalChrom Chromosome) {
    finalChrom.genImage = image.NewRGBA(image.Rect(0, 0, Xmax, Ymax))
    finalChrom.gc = draw2dimg.NewGraphicContext(finalChrom.genImage)

    // splitPoint := int(random(1, GenePoolSize))
    //
    // for i := 0; i < splitPoint; i++ {
    //     finalChrom.triangles[i] = chrom1.triangles[i]
    // }
    //
    // for j := splitPoint; j < GenePoolSize; j++ {
    //     finalChrom.triangles[j] = chrom2.triangles[j]
    // }

    // for i := 0; i < GenePoolSize; i++ {
    //     finalChrom.triangles[i] = chrom1.triangles[i]
    //     finalChrom.triangles[i].color = chrom12.triangles[i].color
    // }

    return
}

func selectParent(currentGen [PopSize]Chromosome) (Chromosome){
    parent1 := currentGen[int(random(0, PopSize))]
    parent2 := currentGen[int(random(0, PopSize))]
    parent3 := currentGen[int(random(0, PopSize))]
    //fmt.Println("Fittest: " + strconv.FormatInt(int64(fittestParent.fitness), 10))

    return findFittest(parent1, parent2, parent3)
}

func exportFittestChromosome(currentGen [PopSize]Chromosome) {
    fittest := currentGen[0]
    for i := 1; i < PopSize; i++ {
        if currentGen[i].fitness < fittest.fitness {
            fittest = currentGen[i]
        }
    }
    fmt.Println("Fittest: " + strconv.FormatInt(int64(fittest.fitness), 10) + "\n")

    filename := "/home/npzd/Pictures/generated/00/image" + strconv.FormatInt(time.Now().Unix(), 10) + ".png"
    draw2dimg.SaveToPngFile(filename, fittest.genImage)
}

func findFittest(parent1, parent2, parent3 Chromosome) (Chromosome) {
    if parent1.fitness < parent2.fitness {
        if parent1.fitness < parent3.fitness {
            return parent1
        } else {
            return parent3
        }
    } else {
        if parent2.fitness < parent3.fitness {
            return parent2
        } else {
            return parent3
        }
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
