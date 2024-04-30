package main

import (
	"fmt"
	"math"
	"math/rand"
)

/*
This code defines a simple neural network with an input layer, a hidden layer, and an output layer.
It uses the sigmoid function as the activation function. The initializeNetwork function creates
the network with random weights and biases. The forwardPass function takes an input vector and
calculates the output of the network.

Note: This is a very basic example and doesn't include functionalities like backpropagation or
training the network. It's meant to give you a starting point for understanding the core
structure of a neural network in Go.
*/

// Neuron defines a single neuron in the network
type Neuron struct {
	weights []float64 // Weights for each input connection
	bias    float64   // Bias value
	output  float64   // Output value of the neuron
}

// Network defines the neural network structure
type Network struct {
	inputLayer  []Neuron // Layer of input neurons
	hiddenLayer []Neuron // Layer of hidden neurons
	outputLayer []Neuron // Layer of output neurons
}

// sigmoid function as activation function
func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

// initializeNetwork creates a network with specified sizes
func initializeNetwork(inputSize, hiddenSize, outputSize int) *Network {
	network := &Network{
		inputLayer:  make([]Neuron, inputSize),
		hiddenLayer: make([]Neuron, hiddenSize),
		outputLayer: make([]Neuron, outputSize),
	}

	// Initialize weights and biases with random values
	for i := range network.inputLayer {
		network.inputLayer[i].weights = make([]float64, hiddenSize)
		for j := range network.inputLayer[i].weights {
			network.inputLayer[i].weights[j] = rand.Float64()
		}
	}

	for i := range network.hiddenLayer {
		network.hiddenLayer[i].weights = make([]float64, outputSize)
		for j := range network.hiddenLayer[i].weights {
			network.hiddenLayer[i].weights[j] = rand.Float64()
		}
	}

	return network
}

// forwardPass calculates the output of the network for a given input
func (n *Network) forwardPass(inputs []float64) []float64 {
	// Calculate hidden layer outputs
	for i := range n.hiddenLayer {
		var sum float64
		for j := range n.inputLayer {
			sum += inputs[j] * n.inputLayer[j].weights[i]
		}
		sum += n.hiddenLayer[i].bias
		n.hiddenLayer[i].output = sigmoid(sum)
	}

	// Calculate output layer outputs
	outputs := make([]float64, len(n.outputLayer))
	for i := range n.outputLayer {
		var sum float64
		for j := range n.hiddenLayer {
			sum += n.hiddenLayer[j].output * n.hiddenLayer[j].weights[i]
		}
		sum += n.outputLayer[i].bias
		n.outputLayer[i].output = sigmoid(sum)
		outputs[i] = n.outputLayer[i].output
	}

	return outputs
}

func main() {
	// Define network size
	inputSize := 2
	hiddenSize := 3
	outputSize := 1

	// Create a network
	network := initializeNetwork(inputSize, hiddenSize, outputSize)

	// Sample input data
	inputs := []float64{0.1, 0.2}

	// Get the network output
	outputs := network.forwardPass(inputs)

	fmt.Println("Network output:", outputs)
}
