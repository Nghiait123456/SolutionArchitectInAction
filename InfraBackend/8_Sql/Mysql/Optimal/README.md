- [Optimal Data Type](#optimal_data_type)
    - [Smaller is usually better](#smaller_is_usually_better)

## Optimal Data Type <a name="optimal_data_type"></a>

## Smaller is usually better <a name="smaller_is_usually_better"></a>

Trong các task nghiệp vụ, việc đầu tiên cần đảm bảo là data type phải đáp ứng được nghiệp vụ. Sau đó, một phần có ý
nghĩa khá quan trọng khi data base phing to và giúp tăng tốc câu query, giảm size câu query là chon được datatype nhỏ
hơn. </br>

Yếu tố khá quan trọng để chon được data type phù hợp là hiểu data type mysql và hiểu đặc thù data của bạn. Tôi phân tích
cụ thể trong các mục dưới. </br>

## Avoid NULL if possible <a name="avoid_null_if_possible"></a>

Nên hạn chế dùng NULL tối đa có thể, chỉ dùng khi cần thiết. Dưới đây, tôi sẽ phân tích cụ thể lý do và các trường hợp:
Lý do nên hạn chế dùng NULL:

1) Khó cho việc tối ưu câu query, nó làm việc lập chỉ mục, tối ưu chỉ mục. Một cột
   NullAble sử dụng nhiều dung lượng hơn và yeu cầu xử lý đặc biệt trong mysql. Khi lập chỉ mục, với cột NullAbe, nó yêu cầu thêm một byte và gây ra chỉ mục cố định.

2) Các phép so sánh với Null cần một xử lý đặc biệt, các phép tính toán là không khả thi. Để so sánh với NULL cần xử dụng IS_NULL, IS_NOT_NULL(), các phép so sánh khác không có giá trị. Tất cả các phép so sánh khác với NULL luôn trả về Fail.
    EX: SELECT * FROM my_table WHERE phone = NULL; allway return emppty mặc dù Phone có default NULL.
    When using DISTINCT, GROUP BY, or ORDER BY, all NULL values are regarded as equal.

   When using ORDER BY, NULL values are presented first, or last if you specify DESC to sort in descending order.
   Aggregate (group) functions such as COUNT(), MIN(), and SUM() ignore NULL values. The exception to this is COUNT(*), which counts rows and not individual column values. For example, the following statement produces two counts.

Lý do nên dùng:  
1) Mysql lưu Null với 1 byte dữ liệu, là điểm cộng khi sử dụng Nullable


Các trường hợp không dùng:

Hầu hết các trường hợp có data rõ ràng và có default data rõ ràng sẽ hạn chế dùng default null:
ex: số tiền trong kho, cân nặng của sản phẩm, status của một trạng thái,....
các data liên quan chặt chẽ đến hàm cần index, thống kê, tổng hợp thường sẽ không đề default null.

Các trường hơp dùng:
Thường chỉ dùng khi data default là NULL chứ không được là giá trị khác. 



