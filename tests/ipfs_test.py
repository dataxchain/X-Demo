import ipfsapi

api = ipfsapi.connect('ipfs-ipfs', 5001)

res = api.add('ipfs_test.txt')
print(res)

content = api.cat(res['Hash'])
print(content)