Welcome to the Exquisite Corpse Image Cropping Utility!
====================================================

Rules of the Game
-----------------

This is a collaborative art game inspired by the Exquisite Corpse drawing technique, developed in the early 20th century by the surreallists. In this game, participants take turns creating a seamless image by cropping off the top or bottom portion of an image and passing it to the next person. Each participant then adds their creative touch to the cropped portion, resulting in a unique and combined artwork.

### Desired Outcome

The goal of the Exquisite Corpse is to create a collection of interconnected images that flow seamlessly from one to another. Each participant contributes by cropping either the top or bottom of an image and then handing it to the next person to continue the chain. The final outcome is an assemblage of diverse artworks that seamlessly fit together.

How to Play
-----------

To play the Exquisite Corpse Image Cropping Game:

1.  Run the provided Python or Go script with the image you want to crop and the desired cropping parameters.
2.  Crop either the top or bottom portion of the image, as specified.
3.  Save the cropped portion as a separate image file.
4.  Draw or Outpaint your contribution, and use the utility to crop off either the top or the bottom, whichever side does not contain the image crop you recieved.
5.  Pass the cropped image to the next participant, who will add their creative artwork to it.
6.  Repeat the process until all participants have contributed.
7.  Combine the final images to reveal the collaborative masterpiece.

### Usage

This script is available in Python and Golang variants, whichever you're more confortable with. The syntax to run the script is as follows:

    python corpse.py <image_path> <top_or_bottom> <pixels>

or

    go run corpse.go <image_path> <top_or_bottom> <pixels>

**Note:** Replace <image\_path>, <top\_or\_bottom>, and <pixels> with the appropriate values. I've supplied a throwaway test.png image for you to play with.

### Examples

To crop the top 300 pixels of an image:

    python corpse.py image.jpg top 300

To crop the bottom 200 pixels of an image:

    python corpse.py image.jpg bottom 200

It's fun, it's creative, it's... art history! 

Have fun playing the Exquisite Corpse with your friends and creating a collaborative artwork that seamlessly blends your creativity with others!
