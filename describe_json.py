#!/usr/bin/env python

import json

FILE_NAME = 'test_data.json'
#OUPUT_FILE = 'out.json'

def describe_dict(a_dict):
    out_dict = dict()
    for k in a_dict.keys():
        if type(a_dict[k]) is list:
            out_dict[k] = '[]{}'.format(type(a_dict[k][0]).__name__ if list else '')
        elif type(a_dict[k]) is dict:
            out_dict[k] = describe_dict(a_dict[k])
        else:
            out_dict[k] = type(a_dict[k]).__name__
    return out_dict

def main():
    #with open(FILE_NAME, 'r') as data_file, open(OUPUT_FILE, 'w') as out_file:
    with open(FILE_NAME, 'r') as data_file:
        json_data = json.load(data_file)
        print(describe_dict(json_data[0]))

if __name__ == '__main__':
    main()
