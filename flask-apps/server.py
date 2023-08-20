from flask import Flask
import time

app = Flask(__name__)

@app.route("/ping")
def ping():
    return "pong"

@app.route("/hello_world")
def hello_world():
    return "<h1>Hello!!! How are you?</h1>"

@app.route("/new_thing")
def new_thing():
    return "<h1>NICE NEW THING</h1>"

@app.route("/process_wait")
def process_wait():
    time.sleep(5)
    return "<h1>Just woke up</h1>"

@app.route("/json_payload")
def get_json_payload():
    time.sleep(2)
    return {
            "mew": "noice"
    }
