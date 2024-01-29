package examples

import "github.com/lukasbm/wandb-go"

func main() {
	run := wandb.init("lukasbm", "test")
	defer wandb.Finish()

	config := map[string]any{
		"epochs":        10,
		"batch_size":    32,
		"learning_rate": 0.001,
	}

	wandb.SetConfig(config)

	for epoch := 0; epoch < config["epochs"]; epoch++ {
		// train
		wandb.Log(map[string]any{
			"loss":     0.5,
			"accuracy": 0.9,
		})
	}
}
