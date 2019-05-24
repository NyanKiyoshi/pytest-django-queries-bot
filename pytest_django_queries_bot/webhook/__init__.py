from flask import request

from . import handlers


def handle_request():
    event_type = request.headers.get("X-Github-Event", None)
    handler = getattr(
        handlers, f"handle_{event_type}_event", handlers.DEFAULT_EVENT_HANDLER
    )
    return handler()
