import binascii

str1 = bytearray.fromhex(input())
str2 = bytearray.fromhex(input())
print(binascii.hexlify(bytearray([a ^ b for (a, b) in zip(str1, str2)])))
