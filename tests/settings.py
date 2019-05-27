from pytest_django_queries_bot.settings import *  # noqa

TESTING = True
SERVER_NAME = "localhost.test"  # RFC2606

SECRET_KEY = "test"
GITHUB_SECRET_KEY = None
WTF_CSRF_ENABLED = False

PASSWORD_CONFIG = {"pbkdf2_sha512__default_rounds": 1}

CELERY_BROKER_URL = None
CELERY_RESULT_BACKEND = None
CELERY_TASK_ALWAYS_EAGER = True
