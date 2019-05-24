COPY requirements*.txt /app/

WORKDIR /app
RUN pip install -r requirements.txt -r requirements_dev.txt
