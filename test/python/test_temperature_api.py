import requests

base_url = "http://localhost:8000"


def make_request(param) -> requests.Response:
    url = f"{base_url}/{param}"
    return requests.request(method="get", url=url)


def test_happy_path_one():
    response = make_request(1811)
    assert response.status_code == 200

    print(f"\n{response.text}")
    assert "1.811000E+00 kK" in response.text
    assert "1.537850e+03 Â°C" in response.text
    assert "1.600221E+03 nm" in response.text


def test_happy_path_decimal():
    response = make_request()
    assert response.status_code == 200

    print(f"\n{response.text}")
