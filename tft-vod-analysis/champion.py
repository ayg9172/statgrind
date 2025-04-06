from dataclasses import dataclass, field
from typing import List
import json
import os
from dataclasses import asdict

@dataclass
class Champion:
    name: str
    traits: List[str] = field(default_factory=list)
    cost: int = 0


def champion_to_json(champion):
    """Convert a Champion dataclass instance to a JSON string."""
    return json.dumps(asdict(champion), indent=4)

def champion_from_json(json_str):
    """Convert a JSON string into a Champion dataclass instance."""
    data = json.loads(json_str)  # Convert JSON string to dictionary
    return Champion(**data)  # Unpack dictionary into Champion

def save_champion(champion):
    """Save a Champion instance as a JSON file in the 'data' folder."""
    os.makedirs("data", exist_ok=True)  # Ensure 'data' folder exists

    file_path = f"data/{champion.name.lower()}.json"
    with open(file_path, "w", encoding="utf-8") as file:
        json.dump(asdict(champion), file, indent=4)

def save_champion(champion):
    """Save a Champion instance as a JSON file in the 'data' folder."""
    os.makedirs("data", exist_ok=True)  # Ensure 'data' folder exists

    file_path = f"data/{champion.name.lower()}.json"
    with open(file_path, "w", encoding="utf-8") as file:
        json.dump(asdict(champion), file, indent=4)