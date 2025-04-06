import re


def previous_stage(stage):
    x, y = map(int, stage.split('-'))

    if x == 1 and y == 1:
        return "1-1"  # Stage 1-1 is the lowest possible stage

    min_y = 1  # Minimum y value
    max_y = 4 if x == 1 else 7  # Stage 1 has max y = 4, others have max y = 7

    if y > min_y:
        return f"{x}-{y-1}"
    else:
        return f"{x-1}-{max_y}"  # Go to the last round of the previous stage


def next_stage(stage):
    x, y = map(int, stage.split('-'))
    
    max_y = 4 if x == 1 else 7  # Stage 1 has max y = 4, others have max y = 7
    
    if y < max_y:
        return f"{x}-{y+1}"
    else:
        return f"{x+1}-1"


def is_stage(st):
    return bool(re.fullmatch(r"\d+-\d+", st))

def less_than(stage1, stage2):
    x1, y1 = map(int, stage1.split('-'))
    x2, y2 = map(int, stage2.split('-'))
    return (x1, y1) < (x2, y2)

def greater_than(stage1, stage2):
    x1, y1 = map(int, stage1.split('-'))
    x2, y2 = map(int, stage2.split('-'))
    return (x1, y1) > (x2, y2)

def stage_of(stage):
    try:
        x, y = map(int, stage.split('-'))
        return x

    except:
        return -1
    
def round_of(stage):
    try:
        x, y = map(int, stage.split('-'))
        return y

    except:
        return -1
    

CAROUSEL = 4
NEUTRALS = 7

def get_minimum_stage_seconds(stage):
    
    x, y = map(int, stage.split('-'))

    if x == 1:
        return 5

    if y == CAROUSEL:
        return 10
    
    if y == NEUTRALS:
        return 30
    
    return 45



MAX_ROUND_LENGTH = 60 * 3

def get_max_stage_key(stages):
    """
    Returns the maximum key from a dictionary where keys are in the format 'x-y'.

    Parameters:
        stages (dict): Dictionary with keys formatted as 'x-y'.

    Returns:
        str: The maximum stage key.
    """
    return max(stages.keys(), key=lambda k: tuple(map(int, k.split('-'))))