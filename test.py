#!/usr/bin/env python3

import os
import subprocess
import yaml


test_index = 0
main_path = './build/uhash'
config_path = './test.yaml'
test_config = yaml.load(open(config_path, 'r'), Loader=yaml.FullLoader)
test_cases = test_config['cases']

class col:
    HEADER = '\033[95m'
    OKBLUE = '\033[94m'
    OKCYAN = '\033[96m'
    OKGREEN = '\033[92m'
    WARNING = '\033[93m'
    FAIL = '\033[91m'
    ENDC = '\033[0m'
    BOLD = '\033[1m'
    UNDERLINE = '\033[4m'

def run_test(params):
    name = params['name']
    cmd = params['cmd']
    output = params['out']
    status = params['code']    
    
    global test_index
    test_index += 1
    print(f'{test_index}/{len(test_cases)}: {name}', end='')
    proc = subprocess.run(f'{main_path} {cmd}', stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True, shell=True)

    # Check status
    if status != None:
        assert proc.returncode == status, \
            f'\n   Status mismatch (expected: {status}, got: {proc.returncode})'
        
    # Check output
    if output != None:
        assert proc.stdout == output or proc.stdout == output + '\n', \
        f'\n   Output mismatch (expected: {output}, got: {proc.stdout})'
        
    print(f' - [ {col.OKGREEN}PASS{col.ENDC} ]')


def run_tests():
    success = 0
    for test in test_cases:
        try:
            run_test(test)
            success += 1
        except AssertionError as e:
            print(f' - [ {col.FAIL}FAIL{col.ENDC} ] {e}')
    return success
    
    
if __name__ == '__main__':
    print('Building...')
    res = os.system(f'make build && chmod +x {main_path}')
    if res != 0:
        print('Build failed with status code:', res)
        exit(1)
    print('Build complete')
    
    success = run_tests()
    all = len(test_cases)
    if success == all:
        print(f'{col.OKGREEN}ALL {success}/{all}{col.ENDC} tests passed!')
        exit(0)
        
    print(f'{success}/{all} tests passed!')
    print(f'{col.FAIL}{all-success}/{all}{col.ENDC} tests have errors!')
    exit(all-success)
