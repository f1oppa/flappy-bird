package main

import "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 480, "flappy bird")

	// Global variables

	alive := true;
	playing := false;

	mode1c := rl.Black;
	mode2c := rl.DarkGray;

	option1c := rl.Black;
	option2c := rl.DarkGray;

	deathmenu := 0;
	mode := 0;

	var score float32 = 0;

	birdx := 20;
	var birdy float32 = 0;

	var wall1x float32 = 300;
	wall1y := 0;

	var wall2x float32 = 300;
	wall2y := 280;

	var wall3x float32 = 700;
	wall3y := 0;

	var wall4x float32 = 700;
	wall4y := 350;

	// Game loop

	for !rl.WindowShouldClose() {
		if(playing){
			if(alive){

				// Drawing the components

				rl.BeginDrawing();

				rl.ClearBackground(rl.RayWhite);

				rl.DrawRectangle(int32(birdx), int32(birdy), 40, 40, rl.DarkGray);

				rl.DrawRectangle(int32(wall1x), int32(wall1y), 50, 100, rl.Black);
				rl.DrawRectangle(int32(wall2x), int32(wall2y), 50, 200, rl.Black);

				rl.DrawRectangle(int32(wall3x), int32(wall3y), 50, 200, rl.Black);
				rl.DrawRectangle(int32(wall4x), int32(wall4y), 50, 180, rl.Black);

				rl.EndDrawing();

				// Handle flying or falling

				if(rl.IsKeyDown(32)){
					for i := 0; i < 1000; i++ {
						birdy -= 0.5 * rl.GetFrameTime();
					}
				} else {
					birdy += 200 * rl.GetFrameTime();
				}

				// Moving the walls

				if(mode == 0) {
					wall1x -= 100 * rl.GetFrameTime();
					wall2x -= 100 * rl.GetFrameTime();
	
					wall3x -= 100 * rl.GetFrameTime();
					wall4x -= 100 * rl.GetFrameTime();
				} else {
					wall1x -= 100 * (score + 1) * rl.GetFrameTime();
					wall2x -= 100 * (score + 1) * rl.GetFrameTime();

					wall3x -= 100 * (score + 1) * rl.GetFrameTime();
					wall4x -= 100 * (score + 1) * rl.GetFrameTime();
				}

				if (wall1x < -50 && wall2x < -50) {
					wall1x = 810;
					wall2x = 810;
				}

				if (wall3x < -50 && wall4x < -50) {
					wall3x = 810;
					wall4x = 810;
				}

				// Handling death events

				if (birdy > 490) {
					alive = false;

					deathmenu = 0;
					option1c = rl.Black;
					option2c = rl.DarkGray;

					mode = 0;
					mode1c = rl.Black;
					mode2c = rl.DarkGray;

				}

				if(birdy < 0) {
					birdy = 0;
				}

				if(wall1x < 40 && wall2x < 40) {
					if(birdy < 50 || birdy > 280) {
						alive = false;

						deathmenu = 0;
						option1c = rl.Black;
						option2c = rl.DarkGray;

						mode = 0;
						mode1c = rl.Black;
						mode2c = rl.DarkGray;

					} else {
						score = score + 0.001;
					}
				}

				if (wall3x < 40 && wall4x < 40) {
					if (birdy < 200 || birdy > 350) {
						alive = false;

						deathmenu = 0;
						option1c = rl.Black;
						option2c = rl.DarkGray;

						mode = 0;
						mode1c = rl.Black;
						mode2c = rl.DarkGray;

					} else {
						score = score + 0.001;
					}
				}
			}
		}

		// Death screen

		if(!alive) {
			score = 0;

			rl.BeginDrawing();
			rl.ClearBackground(rl.RayWhite);

			rl.DrawText("game over", int32(rl.GetScreenWidth()/2-int(rl.MeasureText("game over", 30))/2), 180, 30, rl.DarkBrown);
			rl.DrawText("play again", int32(rl.GetScreenWidth()/2-int(rl.MeasureText("play again", 20))/2), 225, 20, option1c);
			rl.DrawText("main menu", int32(rl.GetScreenWidth()/2-int(rl.MeasureText("main menu", 20))/2), 250, 20, option2c);

			rl.EndDrawing();

			if(rl.IsKeyPressed(rl.KeyUp) || rl.IsKeyPressed(rl.KeyDown)) {
				if(deathmenu == 0) {
					deathmenu = 1;
					option1c = rl.DarkGray;
					option2c = rl.Black;
				} else {
					deathmenu = 0;
					option1c = rl.Black;
					option2c = rl.DarkGray;
				}
			}

			if(rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyEnter)){

				birdx = 20;
				birdy = 0;

				wall1x = 300;
				wall1y = 0;

				wall2x = 300;
				wall2y = 280;

				wall3x = 700;
				wall3y = 0;

				wall4x = 700;
				wall4y = 350;


				if (deathmenu == 0) {
					alive = true;
				} else {
					playing = false;
					alive = true;
				}
			}
		}

		// Main menu

		if(!playing) {
			rl.BeginDrawing();
			rl.ClearBackground(rl.RayWhite);

			rl.DrawText("flappy bird", int32(rl.GetScreenWidth()/2-int(rl.MeasureText("flappy bird", 30))/2), 180, 30, rl.DarkBrown);

			rl.DrawText("classic mode", int32(rl.GetScreenWidth()/2-int(rl.MeasureText("classic mode", 20))/2), 230, 20, mode1c);
			rl.DrawText("hard mode", int32(rl.GetScreenWidth()/2-int(rl.MeasureText("hard mode", 20))/2), 255, 20, mode2c);

			rl.EndDrawing();

			if(rl.IsKeyPressed(rl.KeyUp) || rl.IsKeyPressed(rl.KeyDown)) {
				if(mode == 0) {
					mode = 1;
					mode1c = rl.DarkGray;
					mode2c = rl.Black;
				} else {
					mode = 0;
					mode1c = rl.Black;
					mode2c = rl.DarkGray;
				}
			}

			if(rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyEnter)) {

				birdx = 20;
				birdy = 0;

				wall1x = 300;
				wall1y = 0;

				wall2x = 300;
				wall2y = 280;

				wall3x = 700;
				wall3y = 0;

				wall4x = 700;
				wall4y = 350;

				playing = true;

			}
		}
	}

	rl.CloseWindow()
}