import csv
import json

csvfile = open('/Users/kedarnag/Documents/verbal_questions.csv', 'r')
jsonfile = open('/Users/kedarnag/Documents/verbal.json', 'w')

fieldnames = ("id","title","option_1","option_2","option_3","option_4","answer")
reader = csv.DictReader( csvfile, fieldnames)
for row in reader:
    json.dump(row, jsonfile)
    jsonfile.write('\n')
