import os
import json

path = 'C:\\Users\\Sean Hatfield\\Desktop\\stac-api\\api-stac-2\\storage\\cop\\metadata\\'
directory = os.fsencode(path)

uploadedUri = "QmdNkShw64QsKs5CokKq1NCrWCMorJRK8DwgMFxeWsrGCe"

for file in os.listdir(directory):
    filename = os.fsdecode(file)
    if filename.endswith('.json'):
        with open(path + filename, 'r+') as f:
            data = json.load(f)
            # currentUri = data['file_url']

            # replacedUri = currentUri.replace('NewUriToReplace', uploadedUri)
            # data['file_url'] = replacedUri

            currentUri = data['name']

            replacedUri = currentUri.replace('Stoned Apes', 'Fed Ape')
            data['name'] = replacedUri

            f.seek(0)
            json.dump(data, f, indent=4)
            # print(currentUri)
            # print(replacedUri)
        # print(filename)
    