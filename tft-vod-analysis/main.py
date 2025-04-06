
import cv2
import os 
import pprint
from analysis import *
from logic import *

def list_video_files(folder="video"):
    """List all MP4 files in the given folder."""
    if not os.path.exists(folder):
        print(f"Folder '{folder}' not found.")
        return []
    
    files = [os.path.abspath(folder + "/" + f) for f in os.listdir(folder)]
    
    if not files:
        print("No MP4 files found in the 'video' folder.")
    
    return files

def choose_video(files):
    """Prompt the user to select a video file."""
    if not files:
        print("No video files in ./videos/")
        return None
    
    print("\nAvailable video files:")
    for i, file in enumerate(files):
        print(f"{i + 1}. {file}")
    
    while True:
        try:
            choice = int(input("\nEnter the number of the video you want to open: ")) - 1
            if 0 <= choice < len(files):
                return files[choice]
            else:
                print("Invalid choice. Please enter a valid number.")
        except ValueError:
            print("Invalid input. Please enter a number.")

def get_video_properties(video):
    """Extract and print video properties."""



    # Get video properties
    width = int(video.get(cv2.CAP_PROP_FRAME_WIDTH))
    height = int(video.get(cv2.CAP_PROP_FRAME_HEIGHT))
    fps = video.get(cv2.CAP_PROP_FPS)
    total_frames = int(video.get(cv2.CAP_PROP_FRAME_COUNT))

    # Calculate length (in seconds)
    length = total_frames / fps if fps > 0 else 0

    # Print results
    print(f"Resolution: {width}x{height}")
    print(f"Frames per second (FPS): {fps}")
    print(f"Total frames: {total_frames}")
    print(f"Video length: {length:.2f} seconds")


ROUND_START_ERROR_TOLERANCE = 120




# [starting_frame, ending_frame)
def analyze_video(video, starting_frame, ending_frame=-1):
    if ending_frame == -1:
        ending_frame = int(video.get(cv2.CAP_PROP_FRAME_COUNT))

    fps = int(video.get(cv2.CAP_PROP_FPS))
    video.set(cv2.CAP_PROP_POS_FRAMES, starting_frame)
    current_frame = starting_frame
    frame_exists, frame = video.read()
    
    games = []


    stages = {}
    previous_stage = UNKNOWN_STAGE

    fast_forward_interval = fps * 30
    unknown_stack_size = 0
    max_unknowns = 5
    unknown_forward = fps * 60 * 5

    previous_known_stage = UNKNOWN_STAGE

    carousel_interval = fps * 10

    detect_first_stage = True


    print("Analyzing Game 1...")

    
    while current_frame < ending_frame:

        # If previous stage was unknown, we have to try detect a new game through checking for first stage
        # (First stage has 4 rounds thus the stage number is positioned differently)
        current_stage = detect_stage(frame, detect_first_stage or previous_stage == UNKNOWN_STAGE)



        if len(stages) > 0 and current_stage != UNKNOWN_STAGE and stage_of(current_stage) < stage_of(get_max_stage_key(stages)):

            print("Timestamp:", frame_to_timestamp(current_frame, fps))

            games.append(stages)
            pprint.pprint(sorted(dict_timestamps(stages, fps).items()))
            stages = {}
            detect_first_stage = True



            print("Analyzing Game", str(len(games) + 1) + "...")


        if current_stage in stages or current_stage == UNKNOWN_STAGE: 


            
            if round_of(previous_stage) == CAROUSEL - 1:
                current_frame = current_frame + carousel_interval
            else:
                current_frame = current_frame + fast_forward_interval

            video.set(cv2.CAP_PROP_POS_FRAMES, current_frame)
            frame_exists, frame = video.read()

            previous_stage = current_stage
            continue



        if greater_than(current_stage, "2-1"):
            detect_first_stage = False

        minimum = max(0, current_frame - fast_forward_interval)
        maximum = current_frame


        while maximum - minimum > ROUND_START_ERROR_TOLERANCE:


            current_frame = minimum + (maximum - minimum) //  2

            video.set(cv2.CAP_PROP_POS_FRAMES, current_frame)
            frame_exists, frame = video.read()

            searched_text = detect_stage(frame, detect_first_stage)

            

            if searched_text == current_stage:
                maximum = current_frame
            else:
                minimum = current_frame + 1


        stages[current_stage] = maximum

        # TODO, Add first time the shop is visible!!!

        if round_of(current_stage) == CAROUSEL - 1:
            current_frame = maximum + carousel_interval
        else:
            current_frame = maximum + fast_forward_interval

        video.set(cv2.CAP_PROP_POS_FRAMES, current_frame)
        frame_exists, frame = video.read()
        previous_stage = current_stage


    if stages != {}:
        games.append(stages)
        pprint.pprint(sorted(dict_timestamps(stages, fps).items()))


    champions = {}

    for stages in games:

        for s in stages:
            current_frame = stages[s]
            end_frame = current_frame
            if next_stage(s) in stages:
                end_frame = stages[next_stage(s)] - ROUND_START_ERROR_TOLERANCE

            video.set(cv2.CAP_PROP_POS_FRAMES, current_frame)
            frame_exists, frame = video.read()
            start_gold = detect_gold(frame)


            start_shop = detect_shop(frame)

            if start_shop is None:
                continue 

            for champ in start_shop:
                if champ == "":
                    continue
                if champ not in champions:
                    champions[champ] = 0
                champions[champ] += 1

            video.set(cv2.CAP_PROP_POS_FRAMES, end_frame)
            frame_exists, frame = video.read()

            end_gold = detect_gold(frame)


            if start_gold != end_gold:
                print(s, start_gold, "->", end_gold)




    pprint.pprint(sorted(champions.items()))




def main():


    file_path = choose_video(list_video_files())

    print("Opening: " + file_path)
    video = cv2.VideoCapture(file_path)  # Load video using OpenCV

    if not os.path.exists(file_path):
        print("File not found!")
        exit()

        
    if not video.isOpened():
        print("Could not open video")
        exit()





    get_video_properties(video)


    games = {}

    #while current_f

    frame_count = int(video.get(cv2.CAP_PROP_FRAME_COUNT)) - 1 # -1 for good luck TODO

    analyze_video(video, 0)
    video.release()

if __name__ == "__main__":
    main()
