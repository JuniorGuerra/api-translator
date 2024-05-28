from selenium import webdriver
from selenium.common.exceptions import NoSuchElementException
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.by import By

# local variables
from selenium.webdriver.chrome.options import Options as ChromeOptions

service = Service(r"{}".format("./chromedriver"))

service.start()

chrome_op = ChromeOptions()
chrome_op.add_argument('--no-sandbox')
# chrome_op.add_argument('--headless')  # TODO: Comentar para visualizar
browser = webdriver.Chrome(options=chrome_op, service=service)
browser.implicitly_wait(time_to_wait=5)


def googleTranslate(text, currentLang, targetLang):
    url = f"""https://translate.google.co.in/?
    sl={currentLang}
    &tl={targetLang}
    &text={text}
    """
    print(url)
    try:
        browser.get(url)
        element = browser.find_element(by=By.CLASS_NAME, value='ryNqvb')
        return element.text
    except NoSuchElementException as NoElement:
        print(NoElement)


def deepTraslate(text, currentLang, targetLang):
    url = f"https://www.deepl.com/translator#{currentLang}/{targetLang}/{text}"
    print(url)
    try:
        browser.get(url)
        elements = browser.find_elements(
            By.CSS_SELECTOR, ".--l.--r.sentence_highlight")

        return elements[1].text
    except NoSuchElementException as NoElement:
        print(NoElement)


if __name__ == "__main__":
    # print("\nResult: ", googleTranslate("Hola Mundo", "es", "en"))
    print(deepTraslate("Hola Mundo Cruel", "es", "en"))
