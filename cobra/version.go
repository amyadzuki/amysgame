// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.
package cobra

import (
	"fmt"

	"github.com/amy911/amy911/maths"

	"github.com/amyadzuki/amysgame/data"
	"github.com/amyadzuki/amysgame/human"

	"github.com/g3n/engine/math32"

	"github.com/spf13/cobra"
)

// versionCmd represents the launcher command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version number",
	Long:  `Prints the version number`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version number not available -- in early development")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.

// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.

func init() {
	human.HumanInit = func(b *human.Human) {
		*(b.VboPos.Buffer()) = math32.NewArrayF32(3*len(data.Remap), 3*len(data.Remap))
	}
	human.HumanUpdate = func(b *human.Human, final bool) {
		age, gender, muscle, weight := b.Params()
		var a0, a1, g0, g1, m0, m1, w0, w1 int
		switch {
		case age <= 0.25:
			a0, a1, age = 0, 1, age*4.0
		case age <= 0.5:
			a0, a1, age = 1, 2, maths.Fma(age, 4.0, -1.0)
		case age <= 0.75:
			a0, a1, age = 2, 3, maths.Fma(age, 4.0, -2.0)
		default:
			a0, a1, age = 3, 4, maths.Fma(age, 4.0, -3.0)
		}
		switch {
		case gender <= 0.5:
			g0, g1, gender = 0, 1, gender*2.0
		default:
			g0, g1, gender = 1, 2, maths.Fma(gender, 2.0, -1.0)
		}
		m0, m1, muscle = 0, 1, muscle
		switch {
		case weight <= 0.5:
			w0, w1, weight = 0, 1, weight*2.0
		default:
			w0, w1, weight = 1, 2, maths.Fma(weight, 2.0, -1.0)
		}

		a0 *= 18
		a1 *= 18
		g0 *= 6
		g1 *= 6
		m0 *= 3
		m1 *= 3
		w0 *= 1
		w1 *= 1

		buf := b.VboPos.Buffer()
		for idx, remap := range data.Remap {
			(*buf)[3*idx+0] = updateCoord(data.Float[3*remap+0], age, gender, muscle, weight, a0, a1, g0, g1, m0, m1, w0, w1)
			(*buf)[3*idx+1] = updateCoord(data.Float[3*remap+1], age, gender, muscle, weight, a0, a1, g0, g1, m0, m1, w0, w1)
			(*buf)[3*idx+2] = updateCoord(data.Float[3*remap+2], age, gender, muscle, weight, a0, a1, g0, g1, m0, m1, w0, w1)
		}
		b.VboPos.Update()
	}
}

func updateCoord(data [90]float32, age, gender, muscle, weight float64, a0, a1, g0, g1, m0, m1, w0, w1 int) float32 {
	a0g0m0w := maths.Mix(float64(data[a0+g0+m0+w0]), float64(data[a0+g0+m0+w0]), weight)
	a0g0m1w := maths.Mix(float64(data[a0+g0+m1+w0]), float64(data[a0+g0+m1+w0]), weight)
	a0g0m__ := maths.Mix(a0g0m0w, a0g0m1w, muscle)
	a0g1m0w := maths.Mix(float64(data[a0+g1+m0+w0]), float64(data[a0+g1+m0+w0]), weight)
	a0g1m1w := maths.Mix(float64(data[a0+g1+m1+w0]), float64(data[a0+g1+m1+w0]), weight)
	a0g1m__ := maths.Mix(a0g1m0w, a0g1m1w, muscle)
	a0g____ := maths.Mix(a0g0m__, a0g1m__, gender)
	a1g0m0w := maths.Mix(float64(data[a1+g0+m0+w0]), float64(data[a1+g0+m0+w0]), weight)
	a1g0m1w := maths.Mix(float64(data[a1+g0+m1+w0]), float64(data[a1+g0+m1+w0]), weight)
	a1g0m__ := maths.Mix(a1g0m0w, a1g0m1w, muscle)
	a1g1m0w := maths.Mix(float64(data[a1+g1+m0+w0]), float64(data[a1+g1+m0+w0]), weight)
	a1g1m1w := maths.Mix(float64(data[a1+g1+m1+w0]), float64(data[a1+g1+m1+w0]), weight)
	a1g1m__ := maths.Mix(a1g1m0w, a1g1m1w, muscle)
	a1g____ := maths.Mix(a1g0m__, a1g1m__, gender)
	a______ := maths.Mix(a0g____, a1g____, age)
	return float32(a______)
}

// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.
