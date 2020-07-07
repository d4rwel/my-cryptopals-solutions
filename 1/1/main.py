import base64

hexInput = input()

encodedStr = bytearray.fromhex(hexInput).decode("ascii").encode()
encodedStr = base64.b64encode(encodedStr)

print(encodedStr)
