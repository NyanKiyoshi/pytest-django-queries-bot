from celery import Celery
from flask import Flask

app = Flask(__name__)
celery = Celery(__name__)
