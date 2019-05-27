"""
For references, see:
https://developer.github.com/v3/activity/events/types/#pullrequestevent/."""
from . import tasks


def handle_synchronize_event(payload: dict):
    tasks.create_diff_for_repository.delay()
    return ""
