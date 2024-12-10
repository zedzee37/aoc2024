use std::{fs::File, io::Read};

fn arrange_storage(storage: Vec<usize>) -> Vec<usize> {
    let mut files = Vec::<usize>::new();
    let mut files_remaining: usize= 0;
    let mut blocks_remaining: usize = 0;

    let mut end_ptr = storage.len() - 1;
    if end_ptr & 1 != 0 {
        end_ptr -= 1;
    }

    let mut i = 0;
    while i < storage.len() {
        let size = storage[i];

        if i >= end_ptr {
            break;
        }
        
        if i & 1 == 0 {
            for _ in 0..size {
                files.push(i / 2);
            }
            i += 1;
        } else {
            if files_remaining == 0 {
                let end_size = storage[end_ptr]; 
                files_remaining = end_size;
            }

            let id = end_ptr / 2;
            
            if blocks_remaining == 0 {
                if size == 0 {
                    i += 1;
                    continue;
                }
                blocks_remaining = size;
            }

            files.push(id);

            blocks_remaining -= 1;
            files_remaining -= 1;

            if blocks_remaining == 0 {
                i += 1;
            }

            if files_remaining == 0 {
                end_ptr -= 2;
                continue;
            }
        }
    }

    while files_remaining > 0 {
        files.push(end_ptr / 2);
        files_remaining -= 1;
    }

    return files;
}

fn calculate_checksum(storage: Vec<usize>) -> usize {
    let mut checksum = 0;

    for i in 0..storage.len() {
        let file = storage[i];
        checksum += file * i;
    }
    
    return checksum;
}

fn main() {
    let mut file = File::open("input.txt").unwrap();

    let mut str = String::new();
    file.read_to_string(&mut str).unwrap();
    let buf: Vec<usize> = 
        str.chars()
            .map(|ch| ch.to_digit(10).unwrap() as usize)
            .collect();

    let files = arrange_storage(buf);
    let checksum = calculate_checksum(files);
    println!("{}", checksum);
}
