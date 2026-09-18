open Int

let rec sum v =
        match v with
        | [] -> 0
        | a :: b -> a + sum b


let rec sum_custom v f = 
        match v with
        | [] -> 0
        | a :: b -> f a (sum_custom b f)


let add x y = x + y

let () = print_endline (Int.to_string (sum_custom [10; 20; 30] add))
