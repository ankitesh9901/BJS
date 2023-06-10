def find_subset_sums(nums):
    subset_sums = set()
    subset_sums.add(1)  # Include the empty subset sum

    for num in nums:
        sums_to_add = set()
        for subset_sum in subset_sums:
            new_sum = subset_sum + num
            sums_to_add.add(new_sum)
        subset_sums.update(sums_to_add)

    return subset_sums

# Example usage
nums = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
result = find_subset_sums(nums)

# Print the subset sums in ascending order
sorted_sums = sorted(result)
print(sorted_sums)
