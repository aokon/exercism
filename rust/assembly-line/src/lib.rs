const BASE_PRODUCTION: f64 = 221.0;
const SIXTY_MINUTES: u32 = 60;

pub fn production_rate_per_hour(speed: u8) -> f64 {
    // pattern matching for ranges is only available in nighty version
    let ratio =
        if speed >= 1 && speed < 5 {
            1.0
        } else if speed >= 5 && speed < 9 {
            0.9
        } else if speed >= 9 && speed <= 10 {
           0.77
        } else {
          0.0
        };

    speed as f64 * BASE_PRODUCTION * ratio
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    let hourly_car_num = production_rate_per_hour(speed) as u32;

    hourly_car_num / SIXTY_MINUTES
}
