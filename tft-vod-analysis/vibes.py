import cv2
import numpy as np

def frame_to_timestamp(frame, fps):
    """
    Convert a frame number to a timestamp string (hh:mm:ss).

    Parameters:
        frame (int): The frame number.
        fps (float): Frames per second.

    Returns:
        str: Timestamp formatted as "hh:mm:ss".
    """
    total_seconds = frame / fps
    hours = int(total_seconds // 3600)
    minutes = int((total_seconds % 3600) // 60)
    seconds = int(total_seconds % 60)
    
    return f"{hours:02d}:{minutes:02d}:{seconds:02d}"

def dict_timestamps(frame_dict, fps):
    """
    Convert a dictionary with frame numbers as values into a dictionary with timestamps (hh:mm:ss) as values.

    Parameters:
        frame_dict (dict): Dictionary where each value is a frame number (int or float).
        fps (float): Frames per second.

    Returns:
        dict: A new dictionary with the same keys but with timestamps as values.
    """
    timestamps = {}
    for key, frame in frame_dict.items():
        total_seconds = frame / fps
        hours = int(total_seconds // 3600)
        minutes = int((total_seconds % 3600) // 60)
        seconds = int(total_seconds % 60)
        timestamps[key] = f"{hours:02d}:{minutes:02d}:{seconds:02d}"
    return timestamps