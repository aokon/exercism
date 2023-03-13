pub fn production_rate_per_hour(speed: u8) -> f64 {
    let ratio: f64;
    let num_of_car: f64 = 221.0;

    if speed < 5 {
        ratio = 1.0;
    } else if speed >= 5 && speed < 9 {
        ratio = 0.9;
    } else {
        ratio = 0.77;
    }

    speed as f64 * num_of_car * ratio
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    let hourly_car_num = production_rate_per_hour(speed) as u32;

    hourly_car_num / 60
}
