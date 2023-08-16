from PIL import Image
import os
import sys

def crop_image(image_path, top_or_bottom, pixels):
    img = Image.open(image_path)
    width, height = img.size
    
    if top_or_bottom == 'top':
        cropped_img = img.crop((0, 0, width, pixels))
    elif top_or_bottom == 'bottom':
        cropped_img = img.crop((0, height - pixels, width, height))
    else:
        print("Invalid option for top_or_bottom. Choose either 'top' or 'bottom'.")
        return
    
    output_dir = 'output'
    if not os.path.exists(output_dir):
        os.makedirs(output_dir)
    
    file_name, ext = os.path.splitext(os.path.basename(image_path))
    output_path = os.path.join(output_dir, f"{file_name}_cropped_{top_or_bottom}{ext}")
    
    cropped_img.save(output_path)
    print(f"Cropped image saved as {output_path}")

def main():
    if len(sys.argv) != 4:
        print("Usage: python script.py <image_path> <top_or_bottom> <pixels>")
        return
    
    image_path = sys.argv[1]
    top_or_bottom = sys.argv[2]
    pixels = int(sys.argv[3])
    
    crop_image(image_path, top_or_bottom, pixels)

if __name__ == "__main__":
    main()
