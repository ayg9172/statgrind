
import cv2
import numpy as np
import pytesseract
from logic import *
from vibes import *

# TODO:
# Implement hashing to reduce OCR calls

STAGE_ONE = (820, 0, 60, 39)

STAGE = (760, 0, 60, 39)
RESULTS = (834, 0, 258, 45)
TIME = (1092, 0, 116, 45)
CHAMPIONS = (475, 924, 1012, 156)
GOLD = (950, 882, 55, 30)


def crop_image(image, x, y, width, height):
    return image[y:y+height, x:x+width]  # Crop using NumPy slicing

def display_many(images): 
    for i, image in enumerate(images):
        cv2.imshow("Image " + str(i+1), image)




def extract_components(frame):


    stage = crop_image(frame.copy(), *STAGE)
    results = crop_image(frame.copy(), *RESULTS)
    time = crop_image(frame.copy(), *TIME)
    champions = crop_image(frame.copy(), *CHAMPIONS)

    return (stage, results, time, champions)


def is_uniform(image):
    """Returns True if all pixels in the image are the same color."""
    return np.all(image == image[0, 0])  # Compare all pixels to the first pixel


def extract_champion(frame, i):
    return crop_image(frame, *champion_tuple(i))

def extract_champion_name(frame, i):
    return crop_image(frame, *champion_name_tuple(i))

def champion_name_tuple(i):

    offset_x = i * CHAMPIONS[2] / 5
    offset_y = CHAMPIONS[3] * 3.8 / 5
    width = CHAMPIONS[2] / 6.5
    a = (offset_x, offset_y, width, CHAMPIONS[3] - offset_y)
    return [int(x) for x in a]


def champion_tuple(i):

    offset_x = i * CHAMPIONS[2] / 5
    a = (offset_x, 0, CHAMPIONS[2] / 5, CHAMPIONS[3])
    return [int(x) for x in a]

def dshow(img):
    cv2.imshow("debug", img)
    cv2.waitKey(0)
    cv2.destroyAllWindows()



def visibilitify(img):

    img = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)

    brightness = np.mean(img)
    if brightness < 40:
        _, img = cv2.threshold(img, 50, 255, cv2.THRESH_BINARY_INV)
    else:
        _, img = cv2.threshold(img, 120, 255, cv2.THRESH_BINARY_INV)
    
    return img

def compute_perceptual_hash(image):
    """
    Compute a perceptual hash for an image.
    Assumes the input image is in BGR format (as read by cv2).
    """
    # Convert image to RGB and then to a PIL Image
    pil_img = Image.fromarray(cv2.cvtColor(image, cv2.COLOR_BGR2RGB))
    return imagehash.phash(pil_img)


SIMILARITY_THRESHOLD = 1
UNKNOWN_STAGE = "?-?"
def detect_stage(image, look_for_stage_one=True):

    


    # TODO: Detect if we're picking augments!

    custom_config = r'--oem 3 --psm 7 -c tessedit_char_whitelist=123456789-'


    if look_for_stage_one:
        option1 = visibilitify(crop_image(image, *STAGE_ONE))
        res1 = pytesseract.image_to_string(option1, config=custom_config).strip()
        if is_stage(res1):
            return res1

    option2 = visibilitify(crop_image(image, *STAGE))

    # Extract text
    res2 = pytesseract.image_to_string(option2, config=custom_config).strip()
    
    if is_stage(res2):
        return res2 
    
    return UNKNOWN_STAGE

def detect_gold(image):
    custom_config = r'--oem 3 --psm 7 -c tessedit_char_whitelist=0123456789'

    vis_img = visibilitify(crop_image(image, *GOLD))
    res = pytesseract.image_to_string(vis_img, config=custom_config).strip()
    if res == "":
        return None 
    return int(res)


EMPTY_SHOP = ["", "", "", "", ""]

def detect_shop(src):

    image = crop_image(src, *CHAMPIONS)

    custom_config = r'-c tessedit_char_whitelist=\"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-&\' \" tessedit_user_words_file=/home/artie/Teamfight/retropunk/resources/names.txt'
    #image = visibilitify(image)


    text = pytesseract.image_to_string(image, config=custom_config).strip().lower()
    
    if "sell" in text:
        return None

    shop = ["", "", "", "", ""]

    for i in range(5):
        ch = extract_champion_name(image, i)
        text = pytesseract.image_to_string(ch, config=custom_config).strip().lower()

        shop[i] = text


    return shop

def read_image_from_cv2(image):

    custom_config = "-c tessedit_char_whitelist=ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789&- "

    # Convert to grayscale (improves OCR accuracy)
    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)


    # Extract text
    text = pytesseract.image_to_string(image, config=custom_config)

    return text.strip()  # Remove unnecessary whitespace