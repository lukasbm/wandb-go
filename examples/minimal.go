package main

import "github.com/lukasbm/wandb-go"

func main() {
	wandb.Init("lukasbm", "test")
	defer wandb.Finish()

	epochs := 5

	config := map[string]any{
		"epochs":        epochs,
		"batch_size":    32,
		"learning_rate": 0.001,
	}

	wandb.SetConfig(config)

	for epoch := 0; epoch < epochs; epoch++ {
		// train
		wandb.Log("loss", 0.5)
		wandb.Log("accuracy", 0.9)
	}
}
