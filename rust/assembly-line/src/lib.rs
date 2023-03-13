const BASE_PRODUCTION: f64 = 221.0;
const SIXTY_MINUTES: u32 = 60;

pub fn production_rate_per_hour(speed: u8) -> f64 {
    let ratio: f64;

    // pattern matching for ranges is only available in nighty version
    if speed < 5 {
        ratio = 1.0;
    } else if speed >= 5 && speed < 9 {
        ratio = 0.9;
    } else {
        ratio = 0.77;
    }

    speed as f64 * BASE_PRODUCTION * ratio
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    let hourly_car_num = production_rate_per_hour(speed) as u32;

    hourly_car_num / SIXTY_MINUTES
}
