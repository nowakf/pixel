This is a fork of pixel. The mainline version is better. I forked it to add some functions that, on balance, make the library worse:

GLFW-related:
 - By default stall the main loop unless there is input.
 - Add a 'PostEvent' function to allow callers to cycle the main loop.

Text-related:
 - Three new methods:
	a) Add() - queues up a letter.
	b) Ink() - changes the color of the queue.
	c) Apply() - writes the queue.

The intention here is to allow large amounts of text to be drawn in an efficient manner.

Input-related:
 - Changed the way input works. I had the idea to alter it such that it all comes in through channels. Turns out that's a shit idea. Input gets dropped all over the place. Not content to leave well-designed alone, I changed some stuff around anyway.

 It's uncommented, so whatever.





