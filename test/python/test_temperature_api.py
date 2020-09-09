import requests

base_url = "http://localhost:8000"


def make_get_request(param) -> requests.Response:
    url = f"{base_url}/{param}"
    return requests.request(method="get", url=url)


def test_happy_path_one():
    response = make_get_request(1811)
    assert response.status_code == 200
    text = response.text.encode("iso-8859-1").decode("utf-8")

    print(f"\n{text}")
    assert "1.811000E+00 kK" in text
    assert "1.537850e+03 °C" in text
    assert "1.600221E+03 nm" in text


def test_abs_zero():
    response = make_get_request("0")
    assert response.status_code == 200
    text = response.text.encode("iso-8859-1").decode("utf-8")

    print(f"\n{text}")
    assert "0.000000E+00 K" in text
    assert "-2.731500e+02 °C" in text
    assert "+Inf m" in text