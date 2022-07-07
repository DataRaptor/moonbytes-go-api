from typing import Any
import requests


def get_api_key():
    url: str = "http://localhost:8080"
    response = requests.get(
        "{}/users/get_or_create_alpha_user/{}".format(
            url,
            "aaaaaaaaaaaaaaFzbXRnzWbqdKnJzpUfbBJvELAN19uVr5PJxByK7jYFepr"
        )
    )
    if response.status_code == 200:
        return response.json()

    raise Exception("Could not get alpha api key")


if __name__ == '__main__':

    api_key: str = get_api_key()['api_key']

    api_key = "17b17db7-3a15-4401-8971-48352644b143"

    # Pass through the Gargantuan Reverse Proxy
    url: str = "http://localhost:8080"
    response = requests.get(
        "{}/data/v1/magic_eden/collections".format(url),
        headers={
            'Authorization': api_key
        }
    )
    if response.status_code == 200:
        print(response.json())
    else:
        print(response.content)
