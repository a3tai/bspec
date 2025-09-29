#!/usr/bin/env python3
import sys
import os
sys.path.append('/Users/steverude/Documents/BSpec-Foundations/scripts/generators/json')
from generator import JSONGenerator

def main():
    os.chdir('/Users/steverude/Documents/BSpec-Foundations')
    generator = JSONGenerator('spec/v1', 'sdk/v1/json')
    generator.generate_all()

if __name__ == "__main__":
    main()