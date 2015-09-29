#!/usr/bin/env python3

import argparse
import xlrd
import csv

parser = argparse.ArgumentParser()

parser.add_argument("-wb", "--workbook", help="workbook input filename")
parser.add_argument("-sh", "--sheet", help="sheet name")
parser.add_argument("-out", "--csvout", help="csv output filename")

if __name__ == '__main__':
   args = parser.parse_args()
   wb = xlrd.open_workbook(args.workbook)
   sh = wb.sheet_by_name(args.sheet)
   csv_file = open(args.csvout, 'w')
   wr = csv.writer(csv_file, quoting=csv.QUOTE_MINIMAL, lineterminator='\n')
   for rownum in range(sh.nrows):
       wr.writerow(sh.row_values(rownum))
   csv_file.close()
