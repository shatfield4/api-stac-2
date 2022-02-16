import os
import json

path = 'C:\\Users\\Sean Hatfield\\Desktop\\stac-api\\api-stac-2\\storage\\bacon\\metadata\\'
pathCop = 'C:\\Users\\Sean Hatfield\\Desktop\\stac-api\\api-stac-2\\storage\\cop\\metadata\\'

path = pathCop
directory = os.fsencode(path)

uriToReplace = 'http://128.199.2.204'
apiUri = "https://api.stonedapeclub.com"

description = "In a metaverse nearby cannabis has been made illegal to the point where the only way to obtain some quality kush is to risk it all and grow it yourself. These brave Stoned Apes are passionate in growing and smoking cannabis, so they built secret grow houses and are growing $TOKE daily. In the hopes of having plentiful weed to smoke and cash to stash away. Thousands Stoned Apes and Feds compete on the streets of the metaverse. Dynamic PVP staking game based on wolf.game. On-chain game. Your NFT is also a ticket to our evergrowing community of cannabis & real estate investors and enthusiasts."

for file in os.listdir(directory):
    filename = os.fsdecode(file)
    if filename.endswith('.json'):
        with open(path + filename, 'r+') as f:
            data = json.load(f)
            currentUri = data['file_url']

            replacedUri = currentUri.replace(uriToReplace, apiUri)
            data['file_url'] = replacedUri
            data['description'] = description
            data['image'] = data['file_url']
            del data['file_url']
            del data['custom_fields']


            # currentUri = data['name']

            # replacedUri = currentUri.replace('Stoned Apes', 'Fed Ape')
            # data['name'] = replacedUri

            f.seek(0)
            json.dump(data, f, indent=4)
            # print(currentUri)
            # print(replacedUri)
        # print(filename)
    