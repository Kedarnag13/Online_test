import csv
import json

csvfile = open('/Users/kedarnag/Documents/verbal_questions.csv', 'r')
jsonfile = open('/Users/kedarnag/Documents/verbal_questions.json', 'w')

fieldnames = ("title","option_1","option_2","option_3","option_4","answer")
reader = csv.DictReader( csvfile, fieldnames)
for row in reader:
    json.dump(row, jsonfile)
    jsonfile.write('\n')



# import csv
# import json
#
# csvfilename = '/Users/kedarnag/Documents/verbal_questions.csv'
# jsonfilename = csvfilename.split('.')[0] + '.json'
# csvfile = open(csvfilename, 'r')
# jsonfile = open(jsonfilename, 'w')
# reader = csv.DictReader(csvfile)
#
# fieldnames = ('id', 'title', 'option_1', 'option_2', 'option_3', 'option_4', 'answer')
#
# output = []
#
# for each in reader:
#   row = {}
#   for field in fieldnames:
#     row[field] = each[field]
# output.append(row)
#
# json.dump(output, jsonfile, indent=2, sort_keys=True)
