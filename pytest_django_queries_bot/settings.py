import os.path
from ast import literal_eval
from os import getenv

APP_DIR = os.path.abspath(os.path.dirname(__file__))
PROJECT_ROOT = os.path.abspath(os.path.join(APP_DIR, os.pardir))


def getenv_ast(name, default_value):
    if name in os.environ:
        value = os.environ[name]
        try:
            return literal_eval(value)
        except ValueError as e:
            raise ValueError("{} is an invalid value for {}".format(value, name)) from e
    return default_value


DEBUG = getenv_ast("DEBUG", True)
SECRET_KEY = getenv("SECRET_KEY", "secret")
GITHUB_SECRET_KEY = getenv("GITHUB_SECRET_KEY", None)

if GITHUB_SECRET_KEY:
    GITHUB_SECRET_KEY = bytes(GITHUB_SECRET_KEY, "utf-8")

TEMPLATE_FOLDER = os.path.join(PROJECT_ROOT, "templates")
