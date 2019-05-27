import os.path

import jinja2
from flask import Flask

from .core.extensions import app, celery
from .webhook import dispatch as webhook_handler


def get_config_path():
    config_filename = os.getenv("CONFIG_FILENAME", "settings.py")
    app_dir_path = os.path.realpath(os.path.dirname(__file__))
    config_path = os.path.join(app_dir_path, config_filename)
    return config_path


def create_app() -> Flask:
    # retrieve the settings file to setup flask
    app.config.from_pyfile(get_config_path())
    app.root_path = app.config["PROJECT_ROOT"]

    # configure extensions
    configure_extensions()

    return app


def configure_extensions():
    configure_jinja()
    configure_celery()
    register_routes()


def configure_jinja():
    """This:
     - configures jinja's policy on undefined variables;
     - install any globals needed by the templates;
     - install any extensions used by the templates."""
    # Configure jinja to warn undefined variables.
    # For more information and possibilities, see:
    #     http://jinja.pocoo.org/docs/2.10/api/#jinja2.make_logging_undefined
    undefined_logger = jinja2.make_logging_undefined(app.logger, base=jinja2.Undefined)
    app.jinja_env.undefined = undefined_logger


def configure_celery():
    _celery_prefix = "CELERY_"
    _celery_prefix_len = len(_celery_prefix)
    celeryconf = {
        k[_celery_prefix_len:].lower(): v
        for k, v in app.config.items()
        if k.startswith(_celery_prefix)
    }
    celery.conf.update(celeryconf)


def register_routes():
    app.add_url_rule("/webhook", "webhook", webhook_handler, methods=["POST"])
