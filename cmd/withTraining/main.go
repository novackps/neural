package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/*
This code implements a basic feedforward neural network with sigmoid activation functions,
random weight initialization, and gradient descent for training. It's designed to solve the XOR problem
as a simple demonstration. You can modify it to suit different network architectures,
activation functions, loss functions, and optimization algorithms as needed.
*/

// NeuralNetwork represents a simple feedforward neural network.
type NeuralNetwork struct {
	inputSize    int
	hiddenSize   int
	outputSize   int
	weightsIH    [][]float64 // Weights between input and hidden layer
	weightsHO    [][]float64 // Weights between hidden and output layer
	biasH        []float64   // Bias for hidden layer
	biasO        []float64   // Bias for output layer
	learningRate float64
}

// NewNeuralNetwork creates a new NeuralNetwork instance with random weights and biases.
func NewNeuralNetwork(inputSize, hiddenSize, outputSize int, learningRate float64) *NeuralNetwork {
	rand.Seed(time.Now().UnixNano())

	weightsIH := make([][]float64, inputSize)
	for i := range weightsIH {
		weightsIH[i] = make([]float64, hiddenSize)
		for j := range weightsIH[i] {
			weightsIH[i][j] = rand.Float64() - 0.5 // Random initialization between -0.5 and 0.5
		}
	}

	weightsHO := make([][]float64, hiddenSize)
	for i := range weightsHO {
		weightsHO[i] = make([]float64, outputSize)
		for j := range weightsHO[i] {
			weightsHO[i][j] = rand.Float64() - 0.5
		}
	}

	biasH := make([]float64, hiddenSize)
	for i := range biasH {
		biasH[i] = rand.Float64() - 0.5
	}

	biasO := make([]float64, outputSize)
	for i := range biasO {
		biasO[i] = rand.Float64() - 0.5
	}

	return &NeuralNetwork{
		inputSize:    inputSize,
		hiddenSize:   hiddenSize,
		outputSize:   outputSize,
		weightsIH:    weightsIH,
		weightsHO:    weightsHO,
		biasH:        biasH,
		biasO:        biasO,
		learningRate: learningRate,
	}
}

// Sigmoid activation function.
func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

// Forward propagate input through the network and return the output.
func (nn *NeuralNetwork) forward(input []float64) []float64 {
	hidden := make([]float64, nn.hiddenSize)
	for i := range hidden {
		// Calculate hidden layer values
		sum := 0.0
		for j := 0; j < nn.inputSize; j++ {
			sum += input[j] * nn.weightsIH[j][i]
		}
		hidden[i] = sigmoid(sum + nn.biasH[i])
	}

	output := make([]float64, nn.outputSize)
	for i := range output {
		// Calculate output layer values
		sum := 0.0
		for j := 0; j < nn.hiddenSize; j++ {
			sum += hidden[j] * nn.weightsHO[j][i]
		}
		output[i] = sigmoid(sum + nn.biasO[i])
	}

	return output
}

// Train the neural network using backpropagation.
func (nn *NeuralNetwork) train(input []float64, target []float64) {
	// Forward propagation
	hidden := make([]float64, nn.hiddenSize)
	for i := range hidden {
		sum := 0.0
		for j := 0; j < nn.inputSize; j++ {
			sum += input[j] * nn.weightsIH[j][i]
		}
		hidden[i] = sigmoid(sum + nn.biasH[i])
	}

	output := make([]float64, nn.outputSize)
	for i := range output {
		sum := 0.0
		for j := 0; j < nn.hiddenSize; j++ {
			sum += hidden[j] * nn.weightsHO[j][i]
		}
		output[i] = sigmoid(sum + nn.biasO[i])
	}

	// Backpropagation
	outputErrors := make([]float64, nn.outputSize)
	for i := range output {
		outputErrors[i] = target[i] - output[i]
	}

	hiddenErrors := make([]float64, nn.hiddenSize)
	for i := range hidden {
		sum := 0.0
		for j := 0; j < nn.outputSize; j++ {
			sum += outputErrors[j] * nn.weightsHO[i][j]
		}
		hiddenErrors[i] = sum * hidden[i] * (1 - hidden[i])
	}

	// Update weights and biases
	for i := range nn.weightsHO {
		for j := range nn.weightsHO[i] {
			nn.weightsHO[i][j] += nn.learningRate * outputErrors[j] * hidden[i]
		}
	}

	for i := range nn.weightsIH {
		for j := range nn.weightsIH[i] {
			nn.weightsIH[i][j] += nn.learningRate * hiddenErrors[j] * input[i]
		}
	}

	for i := range nn.biasO {
		nn.biasO[i] += nn.learningRate * outputErrors[i]
	}

	for i := range nn.biasH {
		nn.biasH[i] += nn.learningRate * hiddenErrors[i]
	}
}

func main() {
	// Example usage
	inputSize := 2
	hiddenSize := 3
	outputSize := 1
	learningRate := 0.1

	nn := NewNeuralNetwork(inputSize, hiddenSize, outputSize, learningRate)

	// Training data
	inputs := [][]float64{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	}
	targets := [][]float64{
		{0},
		{1},
		{1},
		{0},
	}

	// Training loop
	for epoch := 0; epoch < 10000; epoch++ {
		for i := range inputs {
			nn.train(inputs[i], targets[i])
		}
	}

	// Test the trained network
	testInput := []float64{1, 0}
	output := nn.forward(testInput)
	fmt.Println("Input:", testInput)
	fmt.Println("Output:", output)
}
